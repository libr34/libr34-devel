async function GetPost() {
    let id = new URL(document.URL).searchParams.get('id')
    console.log(id)

    const resp = await fetch(`https://APIADDR/api/post?id=${id}`, {
        method: 'GET'
    }).then(resp => resp.json()).then(data => {
        let results_pre_str = data
        const results = JSON.parse(results_pre_str)
        console.log(results)
        const source = results.posts[0].file_url
        console.log(source)
        let download = document.getElementById('download')
        download.href = source
        if (source.endsWith(".mp4")) {
            let PostSrc = document.getElementById('sauce')
            PostSrc.src = `${source}`

            let Video = document.getElementById('video-id')
            console.log(Video)
            Video.load()

            //PostSrc.src = source
            //PostSrc.innerHTML = "No videos for now..."

        } else {
            var PostContainer = document.getElementById('PostContainer')

            var Img = document.createElement('img')
            Img.style = " box-shadow: 0.0rem 0.0rem 3rem #ca9ee6;"
            Img.src = source
            PostContainer.appendChild(Img)
        }
    })
}

GetPost()
