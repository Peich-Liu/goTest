// #include <thread>
// #include <vector>
// #include <iostream>

// using namespace std;

// extern "C" {
//     // C 接口函数：启动线程运行算法
//     void RunCppAlgorithm(double* input, int size) {
//         std::thread([input, size]() {
//             double sum = 0;
//             for (int i = 0; i < size; i++) {
//                 sum += input[i];
//             }
//             std::cout << "C++: " << sum << std::endl;
//         }).detach(); // 分离线程，自行管理生命周期
//     }
// }


#include <thread>
#include <vector>
#include <iostream>

using namespace std;


extern "C" void RunCppAlgorithm(double* input, int size);

void RunCppAlgorithm(double* input, int size) {
    std::thread([input, size]() {
        double sum = 0;
        for (int i = 0; i < size; i++) {
            sum += input[i];
        }
        std::cout << "[C++线程] 结果: " << sum << std::endl;
    }).detach();
}
