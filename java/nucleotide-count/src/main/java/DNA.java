import java.util.ArrayList;
import java.util.Collection;
import java.util.Collections;
import java.util.HashMap;
import java.util.TreeSet;

public class DNA {
  private HashMap<Character, Integer> map;

  public DNA(String dna) {
    dna = dna.replaceAll("[^ACGT]", "");
    ArrayList<Character> list = new ArrayList();
    for(char chr : dna.toCharArray()) {
      list.add(new Character(chr));
    }
    TreeSet<Character> set = new TreeSet(list);
    this.map = new HashMap();
    for(char chr : "ACGT".toCharArray()) {
      this.map.put(new Character(chr), 0);
    }
    for(Character chr : set) {
      this.map.put(chr, Collections.frequency(list, chr));
    }
  }

  public int count(char n) {
    if("ACGT".indexOf(n) == -1) {
      throw new java.lang.IllegalArgumentException();
    }
    if(!this.map.containsKey(n)) {
      return 0;
    }
    return map.get(n);
  }

  public Collection nucleotideCounts() {
    return this.map.entrySet();
  }
}
