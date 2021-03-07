module.exports = {
  claimsJwt(req, res) {
    return res.status(200).send({
      status: "claimsJwt",
    });
  },
};
