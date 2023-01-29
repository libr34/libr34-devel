let number_of_pages

//42 there are that many items on one page in default r34

// Get Results
async function getResults() {
    let query = new URL(document.URL).searchParams.get('q')
    console.log(query)

    let page = 0
    page = parseInt(new URLSearchParams(document.URL).get('p'))

    const resp = await fetch(`http://localhost:8080/posts?tags=${query}&p=${page}`, {
        method: 'GET'
    }).then(resp => resp.json()).then(data => {
        let results_pre_str = data
        const results = JSON.parse(results_pre_str)
        //console.log(results.posts.length)
        number_of_pages = Math.round(results.count / 100)

        for (var i = 0; i < page-1; i++) {
            var Atag = document.createElement("a")
            Atag.innerHTML = `${i+1}`
            Atag.href = `./posts.html?q=${query}&p=${page}`

            let pagn = document.getElementById('pagncont')
            pagn.appendChild(Atag)
        }

        for (var i = 0; i < results.posts.length; i++) {
            var NewPostContainer = document.createElement("div")
            var NewAtag = document.createElement("a")
            NewAtag.href = `./post.html?id=${results.posts[i].id}`
            var NewImg = document.createElement("img")
            NewImg.style = "max-width: 300px;"
            NewImg.src = `${results.posts[i].preview_url}`
            var parentDiv = document.getElementById("results")
            NewPostContainer.appendChild(NewAtag)
            NewAtag.appendChild(NewImg)
            parentDiv.appendChild(NewPostContainer)
        }
    })
}
getResults()


let inc = document.getElementById('inc')
let dec = document.getElementById('dec')



const Inc = () => {
    
}

const Dec = () => {

}

inc.addEventListener('click', Inc())
dec.addEventListener('click', Dec())