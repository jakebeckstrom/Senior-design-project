import ScreenshotManager from '../manager/ScreenshotManager.js';
import ResourceManager from '../manager/ResourceManager.js';
import getBrowser from '../utils/browser.js';
import ArtifactManager from '../manager/ArtifactManager.js';
import YaraManager from '../manager/YaraManager.js';
import OCRManager from '../manager/OCRManager.js';
import SafeBrowsingManager from '../manager/SafeBrowsingManager.js';

const processTicket = async ticket => {
  const ticketId = ticket.getID();
  const ticketURL = ticket.getURL();

  const browser = await getBrowser();
  const resourceManager = new ResourceManager();
  const yaraManager = new YaraManager();
  const ocrManager = new OCRManager();
  const ssManager = new ScreenshotManager(await browser.userAgent());
  const safeBrowsingManager = new SafeBrowsingManager();

  console.log(`Starting to process ticket #${ticketId}.`);
  console.log(`URL: ${ticketURL}`);

  // Create a new page in puppeteer:
  const page = await browser.newPage();
  resourceManager.setupResourceCollection(page, ticket);

  await page.goto(ticketURL);

  // Store all artifacts while processing honeyclient, will eventually return to api
  const artifacts = [];

  // process resources, waits for page to finish loading
  const jsArtifacts = await resourceManager.process();
  artifacts.push(...jsArtifacts);

  // needed for safe browsing:
  const urls = resourceManager.getURLs();

  // process screenshots. ocr depends on this, so doing in sync...
  const ssArtifacts = await ssManager.processScreenshots(ticket, page);

  // do next 3 tasks in parallel using Promise.all.
  // ocr, yara, safe browsing API
  const asyncArtifacts = await Promise.all([
    ocrManager.processImages(ssArtifacts),
    yaraManager.setupResourceScan(jsArtifacts, ticket),
    safeBrowsingManager.getMalwareMatches(urls),
  ]);

  artifacts.push(...ssArtifacts); // screenshots
  artifacts.push(...asyncArtifacts[0]); // ocr
  artifacts.push(...asyncArtifacts[1]); // yara

  // store
  ArtifactManager.storeArtifactsForTicket(artifacts);

  await page.close();

  console.log(
    `Finished processing #${ticketId}, ${artifacts.length} artifacts stored.`
  );
  console.log("returning");
  console.log(asyncArtifacts[2]);
  console.log(JSON.stringify(asyncArtifacts[2]));
  console.log(asyncArtifacts[1]);
  console.log(JSON.stringify(asyncArtifacts[1]));
  // Success, return list of paths
  return {
    success: true,
    fileArtifacts: artifacts.map(ArtifactManager.artifactToPath),
    malwareMatches: JSON.stringify(asyncArtifacts[2]),
    yaraMatches: JSON.stringify(asyncArtifacts[1]),
  };
};

export default { processTicket };
