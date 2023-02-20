    // Search bar
    let autoc;

    let search = document.getElementById('search')
    search.addEventListener('input', async function() {
        console.log(search.value)
        if(search.value != undefined && search.value != "") {
        let autoct = document.getElementById('autoct')
        autoct.innerText = `\"${search.value}\"`
        const resp = await fetch(`https://APIADDR/api/autocomplete?q=${search.value}`, {
            method: 'GET',
        })
        if (await resp.ok) {
            autoc = JSON.parse(await resp.json())
            console.log(autoc)
            let autocomp = document.getElementById('autocomplete')
            autocomp.style.visibility = 'visible'
            let autocc = document.getElementById('autocc')
            autocc.innerHTML = ""
            for(let e = 0; e < autoc.length; e++) {
                const newE = document.createElement('a')
                newE.href = `./posts.html?q=${search.value}&p=0`

                const newContent = document.createTextNode(autoc[e].label)
                console.log(autoc[e])

                newE.appendChild(newContent)

                autocc.appendChild(newE)
            }
        }
        } else {
            let autocomp = document.getElementById('autocomplete')
            autocomp.style.visibility = 'hidden'
        }
    })
