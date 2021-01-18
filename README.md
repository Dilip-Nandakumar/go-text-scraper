# text-scraper

Text scraper is used to scan through a given webpage, and display a single consolidated top 10 frequent words and the top 10 frequent word pairs (two words in the same order) along with their frequency. In case the webpage contains hyperlinks, the hyperlinked urls are expanded and the words on these pages are also scanned to come up with the frequent words.

The URL and the maximum number of levels urls are expanded are configurable.

## Steps for running text scraper using docker locally

Run the following commands

```
>> git clone https://github.com/Dilip-Nandakumar/text-scraper.git

>> cd text-scraper

>> make run-docker url=<url> depth=4
```
