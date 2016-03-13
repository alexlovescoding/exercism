public class Pangrams {
  public static boolean isPangram(String str) {
    String ALPHABETS = "abcdefghijklmnopqrstuvwxyz";
    str = str.toLowerCase();
    for (char chr : ALPHABETS.toCharArray()) {
      if (str.indexOf(chr) == -1) {
        return false;
      }
    }
    return true;
  }
}
