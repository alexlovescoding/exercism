import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Set;

public class Etl {
   public Map<String, Integer> transform(Map<Integer, List<String>> old) {
      Set<Integer> keys = old.keySet();
      Map<String, Integer> result = new HashMap<String, Integer>();
      for(Integer key: keys) {
         for(String str: old.get(key)) {
           result.put(str.toLowerCase(), key);
         }
      }
      return result;
   }
}
