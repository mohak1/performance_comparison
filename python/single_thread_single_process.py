import logging
import os
import sys
import time
from typing import List

logging.basicConfig(level=logging.INFO)
log = logging.getLogger()

FILES_DIR = './text_files/smaller_files/'  # dir where log files are present
WORDS_TO_SEARCH = [
    'Session',
    'Warning',
    'Failed',
    'CBS',
    'CSI',
]


def perform_case_insensitive_search(
    words: str,
    line: str,
    tracker: dict
) -> None:
    for word in words:
        tracker[word] += line.lower().count(word.lower())


def validate_path_and_words() -> None:
    log.info('Validating FILES_DIR...')
    if not os.path.exists(FILES_DIR) or not os.path.isdir(FILES_DIR):
        log.error(f'Dir path {FILES_DIR} is not valid')
        raise NotADirectoryError(
            f'The location "{FILES_DIR}" does not exist. Please '
            'check the path and try again.'
        )
    log.info(f'{FILES_DIR} is a valid directory')

    if not WORDS_TO_SEARCH:
        log.error('WORDS_TO_SEARCH list is empty')
        raise ValueError(
            'WORDS_TO_SEARCH cannot be an empty list. Please add some words '
            'to this list and try again.'
        )
    log.info('WORDS_TO_SEARCH list is not empty')

    for word in WORDS_TO_SEARCH:
        if not isinstance(word, str):
            log.error('WORDS_TO_SEARCH list is empty')
            raise TypeError(
                'Entries in WORDS_TO_SEARCH list can only be of type string. '
                f'{word} is not a valid string. Please ensure the values are '
                'strings and try again.'
            )
    log.info('WORDS_TO_SEARCH list contains valid entries')


def process_file(
    file_name: str,
    dir_path: str,
    words_to_search: List[str]
) -> None:
    file_path = os.path.join(dir_path, file_name)
    count_tracker = {i: 0 for i in words_to_search}
    start_time = time.time()
    with open(file_path, 'r') as file:
        for line in file:
            perform_case_insensitive_search(
                words=words_to_search,
                line=line,
                tracker=count_tracker
            )
    duration = time.time() - start_time
    print(
        f'Completed searching {file_name}. Results: {count_tracker}. '
        f'Time taken: {round(duration, 2)} seconds'
    )


def main() -> None:
    try:
        validate_path_and_words()
    except (NotADirectoryError, ValueError, TypeError) as err:
        sys.exit(str(err))

    for file_name in os.listdir(FILES_DIR):
        process_file(
            file_name=file_name,
            dir_path=FILES_DIR,
            words_to_search=WORDS_TO_SEARCH,
        )


if __name__ == '__main__':
    process_start_time = time.time()
    main()
    process_end_time = time.time() - process_start_time
    log.info(f'The process took {round(process_end_time, 2)} seconds')
