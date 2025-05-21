//Can just paste into dev console to test
function GetValue(obj,totalkey) {
    keys = totalkey.split("/")
    pointer = obj
    priorKey = "Top Level Object Does Not Exist"
    for (const key of keys){
        if (!pointer){
            console.log("Bad object! Sub property extracted from key:", priorKey);
            return; 
        }
        if (!(key in pointer)) {
            console.log("Bad key:", key);
            return; 
        }
        pointer = pointer[key] 
        priorKey = key
    }
    console.log(pointer)
}

// Test Cases

//GetValue(null,"a/b/c") 
//GetValue({a:{b:null}},"a/b/c") 
//GetValue({a:{b:{}}},"a/b/c") 
GetValue({a:{b:{c:"Yay"}}},"a/b/c") 
