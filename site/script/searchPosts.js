let number_of_pages

//42 there are that many items on one page in default r34

// Get Results
async function getResults() {
    let query = new URL(document.URL).searchParams.get('q')
    console.log(query)

    let page = 0
    page = parseInt(new URLSearchParams(document.URL).get('p'))
    console.log(page)

    const resp = await fetch(`http://localhost:8080/posts?tags=${query}&p=${page}`, {
        method: 'GET'
    }).then(resp => resp.json()).then(data => {
        let results_pre_str = data
        const results = JSON.parse(results_pre_str)
        //console.log(results.posts.length)
        number_of_pages = Math.ceil(results.count / 42)

        if (page > number_of_pages) {
            window.location.href = `./posts.html?q=${query}&p=${number_of_pages}`
        }


        let pgnm = document.getElementById('pgnm')
        var DecTag = document.createElement("a")
        DecTag.innerHTML = "-"
        if (page <= 0) {
            DecTag.href = `#`
        } else {
            DecTag.href = `./posts.html?q=${query}&p=${page - 1}`
        }
        pgnm.appendChild(DecTag)

        for (var i = 0; i < number_of_pages; i++) {
            var Atag = document.createElement("a")
            Atag.innerHTML = `${i+1}`
            Atag.href = `./posts.html?q=${query}&p=${i}`

            let pagn = document.getElementById('pagncont')

            if (page == i) {
                Atag.style = "margin: 10px; color: #ca9ee6;"
            } else {
                Atag.style = "margin: 10px;"
            }

            pagn.appendChild(Atag)
        }

        /* 
        If the number of pages is less than or equal to the current page then
        we shouldn't allow the user to inc the page number!
        */
        let pgnp = document.getElementById('pgnp')
        var IncTag = document.createElement("a")
        IncTag.innerHTML = "+"
        if (number_of_pages <= page+1) {
            IncTag.href = `#`
        } else {
            IncTag.href = `./posts.html?q=${query}&p=${page + 1}`
        }
        pgnp.appendChild(IncTag)

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