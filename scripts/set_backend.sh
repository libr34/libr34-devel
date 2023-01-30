#!/bin/sh

read -p "Where will your backend be hosted? (r34.example.com/api) <- " api_address

files=("search.js" "searchPosts.js" "videoPlayer.js" "viewPost.js")

for file in "${files[@]}"; do
    sed -i "s/APIADDR/${api_address}/g" "scripts/$file"
done

echo "Done!"