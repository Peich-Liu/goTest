#include <iostream>
#include <vector>
#include "./EDFlib/lib/edflib.h"

extern "C" {

// 加载 EDF 数据并返回指针
double* loadEdf(int* length) {
    std::cout << "Starting EEG processing...\n";

    std::string NewFileName = "c://Document//code4PhD//EEG//Artifact//data//tuh_eeg_artifact//edf//01_tcp_ar//aaaaaaju_s005_t000.edf";      
    if (NewFileName.empty()) {
        std::cerr << "No file specified\n";
        *length = 0;
        return nullptr;
    }

    edf_hdr_struct edf_struct;
    int handle = edfopen_file_readonly(NewFileName.c_str(), &edf_struct, EDFLIB_READ_ALL_ANNOTATIONS);
    if (handle < 0) {
        std::cerr << "Error opening file\n";
        *length = 0;
        return nullptr;
    }

    int sig_idx = 1;
    int total_samples = edf_struct.signalparam[sig_idx].smp_in_datarecord * edf_struct.datarecords_in_file;
    
    if (total_samples <= 0) {
        std::cerr << "No data read\n";
        *length = 0;
        edfclose_file(handle);
        return nullptr;
    }

    std::vector<double> buffer(total_samples);
    int samples_read = edfread_physical_samples(handle, sig_idx, total_samples, buffer.data());

    if (samples_read <= 0) {
        std::cerr << "Failed to read samples\n";
        *length = 0;
        edfclose_file(handle);
        return nullptr;
    }

    // 复制数据到新的动态分配数组（Go 不能直接使用 std::vector）
    double* result = new double[total_samples];
    std::copy(buffer.begin(), buffer.end(), result);

    *length = total_samples; // 传递数据长度
    edfclose_file(handle);
    return result;
}

// 释放动态分配的内存
void freeEdf(double* ptr) {
    delete[] ptr;
}

}
