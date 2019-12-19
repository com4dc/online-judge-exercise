import java.util.Scanner

fun main(args: Array<String>) {
    val scan = Scanner(System.`in`)

    while(true) {
        val nums = scan.nextLine().split(" ")
        val n = Integer.parseInt(nums[0])
        val x = Integer.parseInt(nums[1])

        if (n == 0 && x == 0) {
            break
        }

        println(detect(n, x))
    }
}

fun detect(n: Int, x: Int): Int {
    var count = 0
    for (i in 1..n) {
        for (j in 1..n) {
            for (k in 1..n) {
                if (i != j && j != k && i != k && i + j + k == x && i < j && j < k) {
                    count++
                }
            }
        }
    }
    return count
}
