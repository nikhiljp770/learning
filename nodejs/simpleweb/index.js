
const express = require('express');

const app = express();

app.get('/', (req, res) => {
  res.send('Hi there');
});
app.get('/xyz', (req, res) => {
  
    res.send('Hiiiiii');
  });
  app.get('/new', (req, res) => {
    res.send('new');
  });

  app.get('/reverse', async function(req, res) {
    let string = req.query.string;
  let reversedstring = "";
for (let i = string.length -1 ; i>= 0;i-- )
{
  var str = string.charAt(i)
   reversedstring = reversedstring + str;
}
    res.json(reversedstring);
  });

app.listen(8080, () => {
  console.log('Listening on port 8080');
});