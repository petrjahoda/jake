const name = document.getElementById("name");
const entity = document.getElementById("entity");
const country = document.getElementById("country");
const address = document.getElementById("address");
const city = document.getElementById("city");
const state = document.getElementById("state");
const saveButton = document.getElementById("save-button");
const searchButton = document.getElementById("search-button");
const searchField = document.getElementById("search");
const foundCompany = document.getElementById("found-company");
const table = document.getElementById("table");

UpdateAll();

saveButton.addEventListener("click", (event) => {
    let data = {
        Name: name.value,
        Entity: entity.value,
        Country: country.value,
        Address: address.value,
        City: city.value,
        State: state.value,
    };
    name.value = ""
    entity.value = ""
    country.value = ""
    address.value = ""
    city.value = ""
    state.value = ""
    fetch("/save_one", {
        method: "POST", body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            let result = JSON.parse(data);
            UpdateAll();
            console.log(result)
        });

    }).catch((error) => {
        console.log(error)
    });
})

searchButton.addEventListener("click", (event) => {
    console.log("User input: " + searchField.value)
    let data = {Name: searchField.value,};
    fetch("/search_one", {
        method: "POST", body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            let result = JSON.parse(data);
            if (result["Result"] === "no company found") {
                foundCompany.textContent = "no company found"
                searchField.value = ""
            } else {
                foundCompany.textContent = "company found"
                searchField.value = ""
            }
        });
    }).catch((error) => {
        console.log(error)
    });
})


function UpdateAll() {
    fetch("/get_all", {
        method: "GET",
    }).then((response) => {
        response.text().then(function (data) {
            let result = JSON.parse(data);
            clearTable();
            if (result["Companies"].length > 0) {
                updateTable(result);
            }
        });
    }).catch((error) => {
        console.log(error)
    });
}

function clearTable() {
    let rows = table.rows;
    let i = rows.length;
    while (--i) {
        table.deleteRow(i);
    }
}

function updateTable(result) {
    for (const element of result["Companies"]) {
        table.insertRow().innerHTML =
            "<td>" + element["Name"] + "</td>" +
            "<td>" + element["Entity"] + "</td>" +
            "<td>" + element["Country"] + "</td>";
    }
}


