import java.util.Map;
import java.io.IOException;

public class MindfulParentingSupport {
  
  private Map<String, String> mindfulParentingInformation;
  
  public MindfulParentingSupport() {
    mindfulParentingInformation = new HashMap<String, String>();
  }
  
  public void addMindfulParentingInformation(String type, String info) {
    mindfulParentingInformation.put(type, info);
  }
  
  public Map<String, String> getMindfulParentingInformation() {
    return mindfulParentingInformation;
  }
  
  public String getMindfulParentingInformation(String type) {
    if (!mindfulParentingInformation.containsKey(type)) {
      return null;
    }
    return mindfulParentingInformation.get(type);
  }
  
  public void saveMindfulParentingInformationToFile() throws IOException {
    try {
      FileWriter writer = new FileWriter("mindful_parenting_information.txt");
      
      mindfulParentingInformation.forEach((type, info) -> {
        writer.write(type + ": " + info);
        writer.write("\n");
      });
      
      writer.close();
    } catch (IOException e) {
      throw e;
    }
  }
  
  public void loadMindfulParentingInformationFromFile() throws IOException {
    try {
      FileReader reader = new FileReader("mindful_parenting_information.txt");
      BufferedReader bufReader = new BufferedReader(reader);
      
      String line;
      while ((line = bufReader.readLine()) != null) {
        String[] info = line.split(": ");
        mindfulParentingInformation.put(info[0], info[1]);
      }
      
      bufReader.close();
      reader.close();
    } catch (IOException e) {
      throw e;
    }
  }
  
  public void updateMindfulParentingInformation(String type, String info) {
    mindfulParentingInformation.put(type, info);
  }
  
  public void removeMindfulParentingInformation(String type) {
    if (!mindfulParentingInformation.containsKey(type)) {
      return;
    }
    mindfulParentingInformation.remove(type);
  }
  
  public boolean containsMindfulParentingInformation(String type) {
    return mindfulParentingInformation.containsKey(type);
  }
  
  public static void main(String[] args) {
    MindfulParentingSupport mps = new MindfulParentingSupport();
    mps.addMindfulParentingInformation("Tips for Parents", 
      "Be aware of your own emotions and reactions. \nBe patient and take time to cool down. \nListen to your child and value their perspective. \nStrive for open communication. \nProvide positive feedback and encouragement. \nSet reasonable limits and be consistent with them.");
    mps.addMindfulParentingInformation("Resources",
      "mindful.org \nchildmind.org \nhealthychildren.org \npsychologytoday.com \napa.org \nParenting.org");
    
    try {
      mps.saveMindfulParentingInformationToFile();
      mps.loadMindfulParentingInformationFromFile();
    } catch (IOException e) {
      e.printStackTrace();
    }
  }
  
}