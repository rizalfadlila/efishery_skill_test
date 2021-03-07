const express = require("express");

const app = express();

require("./app/middlewares/middleware.js")(app);
require("./app/router/router.js")(app);

app.listen(process.env.PORT, () => console.log(" run on " + process.env.PORT));
