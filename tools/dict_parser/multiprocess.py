import logging
import multiprocessing

def process_data(data, process_id, progress_queue):
     # 创建日志记录器
    logger = multiprocessing.get_logger()
    logger.setLevel(logging.INFO)

    # 创建文件处理器并设置日志格式
    file_handler = logging.FileHandler(f"process_{process_id}.log")
    formatter = logging.Formatter("%(asctime)s - %(levelname)s - %(message)s")
    file_handler.setFormatter(formatter)

    # 将文件处理器添加到日志记录器
    logger.addHandler(file_handler)

    # 在子进程中记录日志
    logger.info(f"Processing data in process {process_id}")

    # 在这里定义你的处理逻辑

    # 示例日志记录
    logger.info("Log message example")

def process_array_with_progress(data, process_func, num_processes):
    # 计算每个进程应处理的数据数量
    chunk_size = len(data) // num_processes

    # 创建进程池和进程队列
    pool = multiprocessing.Pool(processes=num_processes)
    progress_queue = multiprocessing.Queue()

    # 使用进程池处理数据
    for i in range(num_processes):
        # 计算当前进程应处理的数据范围
        start = i * chunk_size
        end = start + chunk_size if i < num_processes - 1 else None
    
        # 提取当前进程应处理的数据
        chunk_data = data[start:end]

        # 使用进程池异步执行处理函数
        pool.apply_async(process_func, args=(chunk_data, i, progress_queue))

    # 关闭进程池，阻止进一步提交任务
    pool.close()

    # 打印处理进度
    completed_processes = 0
    while completed_processes < num_processes:
        # 阻塞等待进程队列中的进度信息
        progress = progress_queue.get()
        completed_processes += 1
        print(f"Processed by process {progress + 1}/{num_processes}")

    # 等待所有进程完成处理
    pool.join()

    # 所有进程已完成处理
    print("All processes completed.")