# Sort-a-Bytes

Sort-a-Bytes is an webserver that is able to use the system's LaTeX installation to compile .tex files into PDF or render them as PNG.

### Requirements
* Any modern nix box
* latex - to compile .tex files.(See below)
* pdflatex - to convert .tex to pdf
* convert (from imagemagick) - to convert pdfs to pngs


Generally, each distro provides several different texlive packages starting with a barebones version to a "full" installation". The server will run any version. Keep in mind, depending on the version of LaTeX installed, the server will have the ability to compile certain LaTeX files. If you have the extra storage and want to avoid problems just install the "full" LaTeX package associated with your distro.

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
