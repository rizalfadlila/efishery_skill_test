module.exports = {
  claimsJwt(req, res) {
    return res.status(200).send({
      name: req.name,
      phone: req.phone,
      role: req.role,
      timestamp: req.timestamp,
    });
  },
};
