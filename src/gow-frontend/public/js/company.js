function showCompany(str){
    alert(str)
}

function getListCompany(){
    console.log("getcompany")
    var urlCompany = "http://localhost:3000/api/v1/companies"
    $.getJSON(urlCompany, function(listCompany){
        console.log(listCompany[0])
    })
}