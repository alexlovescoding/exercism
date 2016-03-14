import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Anagram {
  private String str;
  public Anagram(String text) {
    str = text.toLowerCase();
  }

  public List<String> match(List<String> list) {
    List<String> result = new ArrayList();
    for (String anagram : list) {
      String tempString = anagram.toLowerCase();
      char[] wrd1 = str.toCharArray();
      char[] wrd2 = tempString.toCharArray();
      Arrays.sort(wrd1);
      Arrays.sort(wrd2);
      if(Arrays.equals(wrd1, wrd2) && tempString.compareTo(str) != 0) {
        result.add(anagram);
      }
    }
    return result;
  }
}
