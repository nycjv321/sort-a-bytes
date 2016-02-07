# Sort-a-Bytes

Sort-a-Bytes is a webserver that is able to use the system's LaTeX installation to compile .tex files into PDF or render them as PNG.

### Requirements
* Any modern *nix box
* latex - to compile .tex files.(See below)
* pdflatex - to convert .tex to pdf
* convert (from imagemagick) - to convert pdfs to pngs


Usually, your OS of choice will provide several different texlive packages to choose from. You will usually find a barebones version that takes about 300-400 MBs of space to a "full" installation that takes about 2-3 gigs of space. Keeping this in mind, the server will have the ability to compile LaTeX files based on the texlive installation package(s) choosen. If you have the extra storage and want to avoid problems, just install the "full" LaTeX package associated with your OS of choice.

### How to use
```sh
# to output pdf
curl -X POST http://server-instance/process -d @example.tex > example.pdf
# to output png
curl -X POST http://server-instance/process?format=png -d @example.tex > example.png
```

### Troubleshooting
This app should run on most UNIX systems. The keyword being "should". The server expects certain commands to exist. If they don't bad things will happen. Check out the following urls to debug the server:
```sh
/sanity-check.json
/sanity-check
```

On my box, sanity-check.json outputs:
```sh
{"convert-installed":true,"latex-installed":true,"pdflatex-installed":true}
```
