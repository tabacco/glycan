Glycan Composition Calculator
=============================

This is a quickly thrown together application for a very specific purpose. Given a weight and three component weights, it'll print out all the combinations of those components that add up to the total weight.

It could almost certainly be more efficient, and there might be bugs and edge cases. On the off chance this is useful to you, that's great! Otherwise, feel free to ignore this. I only hazily understand its purpose myself.

How to use it
-------------
Grab the release for your operating system from the 'Releases' tab and extract the archive.

Run ```./glycan 12345``` to calculate the components adding up to 12345.

By default it'll bracket the input weight by 1, printing results for (weight-1) and (weight+1) as well. You can increase or decrease the bracketing margin with the ```-b``` flag. Passing ```-b 0``` will disable bracketing entirely.
