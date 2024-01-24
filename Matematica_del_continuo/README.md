# Matematica del continuo

L'indice non numera le lezioni in base all'ordine ma in base agli argomenti. Per le dimostrazioni da sapere fate fede al programma completo del vostro anno.

Sotto pdf trovate i file già esportati in pdf. Sotto xopp i file binari (o meglio, xml) che si possono modificare col programma xournal++: https://xournalpp.github.io.

Se avete una cartella con più file .xopp potete convertirli in blocco in pdf e metterli in una sottocartella su bash e simili unix con: `for f in *.xopp; do xournalpp --create-pdf=pdf/${f%.*}.pdf $f; done`
