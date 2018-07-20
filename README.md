Glycan Composition Calculator
=============================

This application calculates which combinations of fucose (146 Da), β-mannose (162 Da), and NAG (203 Da) glycans are needed to explain the difference in glycoprotein mass as a result of N-linked glycosylation from insect cells.

#### Real-life example
* A biochemist conducts a mass spectrometry experiment on a glycoprotein that was purified from Sf9 insect cells. In addition to the predicted native protein mass, a significant peak of +1185 Da is observed. What is it?
* Using the Glycan Composition Calculator and an input of 1185, we learn that the modification comes from 2 NAG, 3 mannose, and 2 fucose residues.
* With this knowledge, the biochemist can select the appropriate strategy for deglycosylating the glycoprotein.

How to use it
-------------
Grab the release for your operating system from the 'Releases' tab and extract the archive.

Given the example data above, run ```./glycan 1185```.

#### Sample Output
```
$ ./glycan 1185

Input Value: 1184 (Adjusted by -1):
Fucose	Mannose	NAG
------	-------	---
2	3	2
7	1	0

Input Value: 1185 (Adjusted by +0):
Fucose	Mannose	NAG
------	-------	---

Input Value: 1186 (Adjusted by +1):
Fucose	Mannose	NAG
------	-------	---

```

By default it'll bracket the input weight by 1, printing results for (weight-1) and (weight+1) as well. This may be useful if your measurements include some margin of error, as the sample one does. You can increase or decrease the bracketing margin with the ```-b``` flag. Passing ```-b 0``` will disable bracketing entirely.

Caveats
-------
It could almost certainly be more efficient, and there might be bugs and edge cases. If you find one, please report it!