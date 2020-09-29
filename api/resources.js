const { documentToHtmlString } = require("@contentful/rich-text-html-renderer");
const client = require("../services/contentful");

module.exports = (_, res) => {
  client.getEntries().then((entries) => {
    res.send(
      entries.items.map((item) => ({
        ...item.fields,
        id: item.sys.id,
        body: documentToHtmlString(item.fields.body),
      }))
    );
  });
};
