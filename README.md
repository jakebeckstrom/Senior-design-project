# Senior-design-project

This repository is a compilation of the separate forked repositories from the Target sponsored senior design project I took part of at the University of Minnesota. I worked with a team of other students to build this app according to specifications given by Target.

## Specification and design
We were tasked with building a web app that used static analysis techniques to analyze a URL. The HoneyClient component utilizes Puppeteer to gather images, scripts and third party information about a specified URL. We used technologies such as Yara static code analysis and Tesseract OCR to analyze scripts and images from a site as well as gathering further information from Google Safe Browsing.

The Frontend contains the User Interface to gather user input and display analysis results.

The API is responsible for handling requests from the Frontend and HoneyClient and storing the informotion in a Postgres database.


## Tech Stack



| Component | Technology |
| ----------- | ----------- |
| Frontend | React, Semantic-UI-React,  |
| API | Golang, Postgres, Docker |
| HoneyClient | Node.js, Express, Puppeteer, Yara, Tesseract OCR, Google Safe Browsing |


Please refer to the README files in each folder for instructions regarding the development environment for each component.

I am incredibly grateful for this opportunity from Target to work with such a talented group of developers and build my skills in full-stack development


