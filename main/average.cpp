#include <vector>
#include <iostream>

extern "C" {
    // 计算平均数
    double compute_average(int* data, int length) {
        if (length == 0) return 0.0;

        std::vector<int> vec(data, data + length);

        // 计算总和
        double sum = 0.0;
        for (int v : vec) {
            sum += v;
        }

        // 输出传入数据，便于调试
        std::cout << "Received data in C++: ";
        for (int v : vec) {
            std::cout << v << " ";
        }
        std::cout << std::endl;

        return sum / length;
    }
}
