const contentful = require("contentful");

module.exports = contentful.createClient({
  accessToken: process.env.NODE_CONTENTFUL_ACCESS_TOKEN,
  space: process.env.NODE_CONTENTFUL_SPACE_ID,
});
