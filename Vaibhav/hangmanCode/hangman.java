import java.util.Arrays;
import java.util.Scanner;


public class Hangman{

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        
        //add some comments in here
        String word = "elephant";
        String hint = "Biggest animal";
        char[] guesses = {'_','_','_','_','_','_','_','_'};
        int chances = 7;

        while(chances>0 && checkArrayContains('_', guesses)){
            boolean matched = false;
            System.out.println("Hint : "+hint);
            System.out.println(Arrays.toString(guesses));
            System.out.println("Chances remaining : "+chances);
            System.out.println("Guess a letter : ");
            char guess = sc.next().charAt(0);

            for(int i=0;i<word.length();i++){
                if(word.charAt(i)==guess && guesses[i]=='_'){
                    guesses[i] = guess;
                    matched = true;
                    break;
                }
            }

            if(matched){
                System.out.println("Letter guessed correctly..");
            }else{
                System.out.println("oops!! wrong letter .. please try again");
                chances--;
            }
            

        }

        if(chances!=0){
            System.out.println("Congratulation you have guessed all charactors correctly..");
        }else{
            System.out.println("You lost all chances please try again");
        }

        sc.close();
    }

    public static boolean checkArrayContains(char a,char[] arr){
        for(char c:arr){
            if(c==a){
                return true;
            }
        }
        return false;
    }

}