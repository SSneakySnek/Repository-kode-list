
public class HelloWorld
{
    public static void main (String[] args) {
        
        //Variable store data, numbers and characters
        int a =5;
        char b = 's';
        long c = 500;
        double d = 3.2;
        
        ------------------------------------------------
        
        //Storing String
        String name = "Thomas";
        
        //How to print out line
        System.out.println("Hello World, this is my first Java line!");
        System.out.println(name.toUpperCase());
        
        ----------------------------------------------------
        
        //Methods
        
        public static void main (String[] args) {
        
        addExclamationPoint("bruh");
        
    }
      
        public static void addExclamationPoint(String s) {
            System.out.println(s + "!");
        
---------------------------------------------------------------
    //return in different classes
    
    public class Brah {
    
    public static String iAmBrah() {
        return "i am a brah";
    }
    
}
 public class HelloWorld {

    public static void main (String[] args) {
        
        Brah a = new Brah();
        String brah =  a.iAmBrah();
        System.out.println(brah);
    }
   
 }
 --------------------------------------------------------------
     
     
