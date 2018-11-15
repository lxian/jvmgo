public class ExceptionTest {
    public static void main(String[] args) {
        try {
            System.out.println("Throwing Exception...");
            throw new Exception("Errorrrrrr");
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }
}
