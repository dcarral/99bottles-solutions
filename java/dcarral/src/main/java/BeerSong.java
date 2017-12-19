public class BeerSong {

    public String verse(int verseNumber) {
        if (verseNumber == 2) {
            return "2 bottles of beer on the wall, 2 bottles of beer.\n" +
                    "Take one down and pass it around, 1 bottle of beer on the wall.\n\n";
        } else if (verseNumber == 1) {
            return "1 bottle of beer on the wall, 1 bottle of beer.\n" +
                    "Take it down and pass it around, no more bottles of beer on the wall.\n\n";
        } else if (verseNumber == 0) {
            return "No more bottles of beer on the wall, no more bottles of beer.\n" +
                    "Go to the store and buy some more, 99 bottles of beer on the wall.\n\n";
        } else {
            return new StringBuilder().append(verseNumber)
                    .append(" bottles of beer on the wall, ")
                    .append(verseNumber)
                    .append(" bottles of beer.\n")
                    .append("Take one down and pass it around, ")
                    .append(verseNumber - 1)
                    .append(" bottles of beer on the wall.\n\n")
                    .toString();
        }

    }

    public String sing(int startingVerseNumber, int endingVerseNumber) {
        StringBuilder verses = new StringBuilder();

        for (int verseNumber = startingVerseNumber; verseNumber >= endingVerseNumber; verseNumber--) {
            String currentVerse = this.verse(verseNumber);
            verses.append(currentVerse);
        }

        return verses.toString();
    }

    public String singSong() {
        return this.sing(99, 0);
    }
}