class BeerSong {

    companion object {

        fun lyrics() = verses(99, 0)

        fun verses(to: Int, from: Int): String {
            return (from..to)
                    .reversed()
                    .map({ verse(it) })
                    .joinToString("")
        }

        fun verse(bottleCount: Int): String {
            return when (bottleCount) {
                0 -> "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n\n"
                1 -> "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n\n"
                2 -> "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n\n"
                else -> "$bottleCount bottles of beer on the wall, $bottleCount bottles of beer.\nTake one down and pass it around, ${bottleCount - 1} bottles of beer on the wall.\n\n"
            }
        }


    }


}