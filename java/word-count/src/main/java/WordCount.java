import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.HashMap;
import java.util.Map;
import java.util.TreeSet;
import java.util.regex.Pattern;

public class WordCount {
  public final Map<String, Integer> phrase(String input) {
    Map<String, Integer> result = new HashMap<String, Integer>();
    Pattern ptr = Pattern.compile("[\\p{Punct}\\s]");
    ArrayList<String> list = new ArrayList(Arrays.asList(ptr.split(input.toLowerCase())));
    for (String word : new TreeSet<String>(list)) {
      if (!word.equals("")) {
        result.put(word, Collections.frequency(list, word));
      }
    }
    return result;
  }
}
