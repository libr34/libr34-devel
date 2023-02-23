# libr34-devel
- It does the funny :3 uwu
### The 34th rule of the internet but the CIA can't shoot your dog UWU!
## How to use:
- Clone the repo
```
git clone https://github.com/libr34/libr34-devel
```
- CD into the repos directory
```
cd libr34-devel
```

### Run your preferred install target:
Make sure you run them in the root of the repository.
```
# make nginx
```
```
# make apache
```
Setting the APIADDR 
- You can go into each script file manually and change APIADDR to {your_domain}/api
- Or edit the files in /var/www/html/libr34/script/ so that you change APIADDR to {your_domain}/api

- In the libr34 directory run:
```
go run .
```
<img alt="GIF of installation" src="https://benike.monster/demo.gif" width="600"/>

## Instances
| URL                        | Country |
|----------------------------|---------|
| https://r34.benike.monster | US      |

## Our devs
- CEO+Sigma Male: benilol98
- Backend dev+Sigma Male: snekatmiskolc
- Frontend dev (kinda)+Sigma Male: namesvin
## TODO / Plans
- Fix pagination issues when there are tons of pages
- Make APIADDR replacement fully automatic
- Make autocomplete add the prompt to the search bar instead of redirecting the user
- Serve everything trough the go backend so that the user doesn't have to touch r34.
