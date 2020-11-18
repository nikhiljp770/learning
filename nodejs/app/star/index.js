const express = require('express');

const app = express();

app.get('/star',async function(req,res){
    let string = req.query.string;
    if(string = "left"); {
        let output = "";
        for(let i = 1; i <= 5; i++ ){
            for(let j = 1; j <= i; j++) {
             output = output + "*";

            }
            output = output + "\n";
        }
        res.json(output)
    }
    });
app.listen(5050, () =>{
    console.log('listening on port 5050');
});