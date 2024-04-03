import java.util.HashMap;

class SqueakyClean {
    static String clean(String identifier) {
        StringBuilder sb = new StringBuilder();

        var ignoreMap = getIgnoredMap();

        for (int i = 0; i < identifier.length(); i++) {

            var letter = identifier.charAt(i);

            if (ignoreMap.containsValue(letter)) {
                continue;
            }

            switch (letter) {
                case '-':
                    var nextLetter = identifier.charAt(i + 1);
                    sb.append(Character.toUpperCase(nextLetter));
                    i++;
                    break;

                case ' ':
                    sb.append('_');
                    break;

                case '0':
                    sb.append('o');
                    break;

                case '1':
                    sb.append('l');
                    break;

                case '3':
                    sb.append('e');
                    break;

                case '4':
                    sb.append('a');
                    break;

                case '7':
                    sb.append('t');
                    break;

                default:
                    sb.append(letter);
                    break;

            }
        }

        return sb.toString();
    }

    private static HashMap<Character, Character> getIgnoredMap() {
       var ignoreMap = new HashMap<Character, Character>();
       ignoreMap.put('!', '!');
       ignoreMap.put('@', '@');
       ignoreMap.put('#', '#');
       ignoreMap.put('$', '$');
       ignoreMap.put('%', '%');
       ignoreMap.put('^', '^');
       ignoreMap.put('&', '&');
       ignoreMap.put('*', '*');
       ignoreMap.put('(', '(');
       ignoreMap.put(')', ')');
       ignoreMap.put('+', '+');
       ignoreMap.put('=', '=');
       ignoreMap.put(',', ',');
       ignoreMap.put('.', '.');
       ignoreMap.put('<', '<');
       ignoreMap.put('>', '>');
       ignoreMap.put(';', ';');
       ignoreMap.put(':', ':');
       ignoreMap.put('¡', '¡');

       return ignoreMap;
    }

}
