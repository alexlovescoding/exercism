import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Anagram {
  private String str;
  private char[] chars;
  public Anagram(String text) {
    str = text.toLowerCase();
    chars = str.toCharArray();
    Arrays.sort(chars);
  }

  public List<String> match(List<String> list) {
    List<String> matches = new ArrayList();
    for (String anagram : list) {
      String tempString = anagram.toLowerCase();
      char[] chars = tempString.toCharArray();
      Arrays.sort(chars);
      if(Arrays.equals(this.chars, chars) && tempString.compareTo(str) != 0) {
        matches.add(anagram);
      }
    }
    return matches;
  }
}
