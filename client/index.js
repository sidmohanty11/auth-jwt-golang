const express = require('express');
const app = express();
const ejsMate = require('ejs-mate');
const port = 4000;
const path = require('path');

app.engine('ejs', ejsMate);
app.set('view engine', 'ejs');
app.set('views');
app.set('views', path.join(__dirname, 'views'));

app.get('/', (req, res) => {
  res.render("index.ejs");
});

app.listen(port, () => {
  console.log(`listening at http://localhost:${port}`);
})