require("dotenv").config();

const express = require("express");
const logger = require("morgan");
const mwCross = require("./cors.js");

module.exports = function (app) {
  app.use(logger("dev"));
  app.use(express.json());
  app.use(
    express.urlencoded({
      extended: true,
    })
  );

  app.use(mwCross());
};
