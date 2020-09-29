const { documentToHtmlString } = require("@contentful/rich-text-html-renderer");
const client = require("../services/contentful");

module.exports = (req, res) => {
  client.getEntry(req.query.entry_id).then((item) => {
    res.send({
      ...item.fields,
      id: item.sys.id,
      body: documentToHtmlString(item.fields.body),
    });
  });
};
