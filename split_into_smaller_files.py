import logging
import os
import sys
import time

logging.basicConfig(level=logging.INFO)
log = logging.getLogger()

SOURCE_FILE_PATH = './text_files/Windows.log'  # big file location
DESTINATION_DIR = './text_files/smaller_files'  # dir to save the files
SMALL_FILE_PREFIX = 'small_file'  # prefix of each smaller file
EXTENSION = '.txt'
NUM_SMALL_FILES = 20  # the number of smaller files


def validate_path_and_values():
    log.info('Validating SOURCE_FILE_PATH...')
    if not os.path.exists(SOURCE_FILE_PATH):
        raise FileNotFoundError(
            f'The location "{SOURCE_FILE_PATH}" does not exist. Please '
            'check the path and try again.'
        )

    if not os.path.isfile(SOURCE_FILE_PATH):
        raise TypeError(
            f'The location "{SOURCE_FILE_PATH}" is not a file. Please '
            'ensure that the path points to a file and try again.'
        )

    log.info(f'{SOURCE_FILE_PATH} is valid file path')
    log.info('Validating DESTINATION_DIR...')
    if not os.path.exists(DESTINATION_DIR) or not os.path.isdir(DESTINATION_DIR):
        log.info(f'{DESTINATION_DIR} doesnt exist, creating a directory')
        os.mkdir(DESTINATION_DIR)
    else:
        log.info(f'{DESTINATION_DIR} is a valid directory')

    log.info('Validating SMALL_FILE_PREFIX...')
    if not isinstance(SMALL_FILE_PREFIX, str) or len(SMALL_FILE_PREFIX) < 1:
        raise ValueError(
            f'The name "{SMALL_FILE_PREFIX}" is not a valid name. Please '
            'ensure the input can be used as a valid file name.'
        )
    log.info(f'{SMALL_FILE_PREFIX} is valid prefix for small files')

    log.info('Validating NUM_SMALL_FILES...')
    if not isinstance(NUM_SMALL_FILES, int) or NUM_SMALL_FILES < 1:
        raise ValueError(
            f'The number "{NUM_SMALL_FILES}" is not valid. Please ensure the '
            'input number is a valid integer greater than 0 and try again.'
        )
    log.info(f'{NUM_SMALL_FILES} is a valid number')


def get_new_small_file_name_and_path(file_num):
    file_name = SMALL_FILE_PREFIX + f'_{file_num}{EXTENSION}'
    file_path = os.path.join(DESTINATION_DIR, file_name)
    return file_name, file_path


def main():
    try:
        log.info('Starting validtion')
        validate_path_and_values()
        log.info('Validtion completed')
    except (
        FileNotFoundError,
        TypeError,
        NotADirectoryError,
        ValueError,
    ) as err:
        sys.exit(err)

    source_file_size = os.path.getsize(SOURCE_FILE_PATH)
    smaller_file_target_size = source_file_size // NUM_SMALL_FILES
    log.info(f'Source file size is {source_file_size} bytes')
    log.info(f'Creating {NUM_SMALL_FILES} smaller files')
    log.info(f'Each small file will be about {smaller_file_target_size} bytes')

    with open(SOURCE_FILE_PATH, 'r') as file:
        file_num = 1
        curr_size = 0
        start_time = time.time()
        small_file_name, small_file_path = get_new_small_file_name_and_path(file_num)
        log.info(f'Starting to create {small_file_path}')
        small_file = open(small_file_path, 'w')

        for line in file:
            if curr_size >= smaller_file_target_size:
                small_file.close()
                duration = time.time() - start_time
                log.info(
                    f'Created {small_file_name} with size {curr_size} bytes in '
                    f'{duration} seconds'
                )
                # resetting vairables; opening new file
                file_num += 1
                curr_size = 0
                start_time = time.time()
                small_file_name, small_file_path = get_new_small_file_name_and_path(file_num)
                log.info(f'Starting to create {small_file_name}')
                small_file = open(small_file_path, 'w')
            else:
                curr_size += len(line.encode('utf-8'))
                small_file.write(line)


if __name__ == '__main__':
    main()
