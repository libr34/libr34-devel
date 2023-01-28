    // Search bar
    let autoc;

    let search = document.getElementById('search')
    search.addEventListener('input', async function() {
        console.log(search.value)
        const resp = await fetch(`http://localhost:8080/autocomplete?q=${search.value}`, {
            method: 'GET',
        })
        console.log(await resp)
        if (await resp.ok) {
            autoc = await resp.json()
            console.log(autoc)
        }
    })
