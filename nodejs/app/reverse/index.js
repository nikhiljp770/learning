const express = require('express');
const app = express();
const fs = require('fs');


app.get('/reverse',async function(req,res){
    //let string = req.query.string;
    fs.readFile("/usr/app/reverse/reverse.txt", 'utf8', function (err,data) {
        if (err) {
          return console.log(err);
        }
        console.log(data);
        data = data.trim();
        let rerversedstring = "";
    for (let i= data.length-1;i>=0;i--)
    {
        var str = data.charAt(i)
    rerversedstring =rerversedstring + str;
    }
    res.json(rerversedstring);
      });
    
});
app.listen(8080, () =>{
    console.log('listening on port 8080');
});