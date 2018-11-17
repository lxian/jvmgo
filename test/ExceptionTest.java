public class ExceptionTest {
    public static void main(String[] args) throws Exception {
        tryException();
    }

    public static void tryException() throws Exception {
        try {
            System.out.println("Throwing Exception...");
            if (false) {
                throw new FooError();
            }
            throw new BarError("Some something went wrong...");
        } catch (FooError e) {
            System.out.println("Caught FooError");
        }
    }
}

class BarError extends Exception {
    public BarError(String message) {
        super(message);
    }
}

class FooError extends Exception {
}