import java.util.Scanner;
import java.util.List;
import java.util.Arrays;

public class Main {
    
    public static void main(String... args) {
        Scanner scanner = new Scanner(System.in);
        
        int count = Integer.parseInt(scanner.nextLine());
        String[] numbers = scanner.nextLine().split(" ");
        
        List<String> list = Arrays.asList(numbers);
        String result = list.stream().reduce((e1, e2) -> e2.concat(" " + e1)).get();
        System.out.println(result);
    }
}
