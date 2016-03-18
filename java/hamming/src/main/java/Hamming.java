public class Hamming {
  public static int compute(String first, String second) {
    if (first.length() != second.length()) {
      throw new java.lang.IllegalArgumentException("Both the strings need
                                                    to be the same length");
    }

    char[] firstChars = first.toCharArray();
    char[] secondChars = second.toCharArray();
    int count = 0;

    for(int i = 0; i < firstChars.length; i++) {
      if(firstChars[i] != secondChars[i]) count++;
    }

    return count;
  }
}
