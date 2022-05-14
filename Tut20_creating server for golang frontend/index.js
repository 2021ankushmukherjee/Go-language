const express = require("express");
const app = express();
port = 4000;


app.use(express.json())
app.use(express.urlencoded({extended: true}))

app.get("/", (req,res) =>{

    res.status(200).json({

        message: "Welcome to my server"
    })
})

app.post("/post", (req,res) =>{

    const myjson = req.body;
    
    res.status(200).send(myjson)
});


app.post("/postform", (req,res)=>{

    res.status(200).send(JSON.stringify(req.body));
});



app.listen(port, ()=>{

    console.log(`server is running on ${port}`);
});