import java.util.Set;

class SqueakyClean {

    private static final Set<Character> specialChars = Set.of('!', '¡', '@', '#', '$', '%', '^', '&', '*', '(', ')', ':', ';', '.', '?');

    /**
     * Cleans the given identifier by applying specific character transformations and removals.
     * <p>
     * This method transforms the input identifier according to the following rules:
     * <ul>
     *     <li>Replaces spaces (' ') with underscores ('_').</li>
     *     <li>Converts leet speak characters to their alphabetical counterparts ('0' to 'o', '1' to 'l', '3' to 'e', '4' to 'a', '7' to 't').</li>
     *     <li>Removes special characters listed in {@code specialChars}.</li>
     *     <li>Transforms kebab-case to camelCase by removing hyphens ('-') and capitalizing the following character.</li>
     * </ul>
     * Note: If a hyphen is at the end of the identifier or is followed by a special character, it is simply removed without further transformation.
     * </p>
     *
     * @param identifier The string identifier to be cleaned.
     * @return A new string representing the cleaned identifier.
     */
    static String clean(String identifier) {
        var sb = new StringBuilder();

        for (int i = 0; i < identifier.length(); i++) {
            var letter = identifier.charAt(i);

            if (specialChars.contains(letter)) {
                continue;
            }

            switch (letter) {

                case ' ' -> sb.append('_');

                case '-' -> {
                    if (i + 1 < identifier.length()) {
                        sb.append(Character.toUpperCase(identifier.charAt(i + 1)));
                        i++;
                    }
                }

                case '0' -> sb.append('o');

                case '1' -> sb.append('l');

                case '3' -> sb.append('e');

                case '4' -> sb.append('a');

                case '7' -> sb.append('t');

                default -> sb.append(letter);
            }
        }

        return sb.toString();
    }

}
