#!/bin/bash

read -p "Where will your backend be hosted? (r34.example.com/api) <- " api_address

sed -i s/APIADDR/$api_address/g site/script/search.js
sed -i s/APIADDR/$api_address/g site/script/searchPosts.js
sed -i s/APIADDR/$api_address/g site/script/videoPlayer.js
sed -i s/APIADDR/$api_address/g site/script/viewPost.js

echo "Done!"