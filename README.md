# Persona
With the barrage of spam/scam calls and emails I decided I needed a fast way to produce Credit Card, SSN, and a name with which to play with these people.

## To Run
First you need golang installed<br>
You can run with "go run ." or you can build for your environment<br>
to run help:<br>
./persona or ./persona -h<br>
Welcome to help<br>
Command line usage is<br>
"go run . -t visa -q 1"<br>
Card Types:<br>
map[amex:{American Express [37 34] 15 4} dci:{Diners Club International [36 38] 16 3} dcu:{Diners Club USA & Canada [54] 16 3} discover:{Discover [6011 65 644 645 646 647 648 649] 16 3} jcb:{Japan Credit Bureau [3528 3529 3530 3531 3532 3533 3534 3535] 16 3} mae:{Maestro [5018 5020 5038 5893 6304 6759 6761 6762 6763] 16 3} maeuk:{Maestro UK [6759 676770 676774] 16 3} mc:{Mastercard [51 52 53 54 55] 16 3} visa:{Visa [4] 16 3}]<br>

To generate a persona:<br>
either "go run . -t visa -q 1"<br>
or<br>
"./persona -t visa -q 1"<br>

To save to a file just add "-w" and it will output to a file

Written in GO and will have a GUI to go with it soon.<br>

Future updates...<br>
-save persona for reference<br>
-gui

## Contributing

Want to contribute? Awesome! The most basic way to show your support is to star
the project, or to raise issues. You can also support this project by
[**becoming a sponsor on GitHub**](https://github.com/sponsors/fizzywhizbang)

Thanks again for your support, it is much appreciated! :pray: