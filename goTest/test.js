'use'

var arr = [{"aa":"AA"},{"aa":"BB"},{"aa":"CC"}]
var newstr = arr.map(v=>({
    "11":v.aa
}))
console.log(newstr)

var newstrr = arr.map(v=>{
    if(v.aa == "aa"){
        var result = {};
        var key = "aa";
        var value = v.aa;
        result[key] = value;
        return result;
    }else{
        var result = {};
        var key = "ww";
        var value = v.aa;
        result[key] = value;
        return result;
    }
})
console.log(newstrr)