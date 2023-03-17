from setuptools import setup, find_packages

VERSION = '0.0.1'
DESCRIPTION = 'Python Package for Every shillings proto'
LONG_DESCRIPTION = 'Python Package for Every shillings proto'

# Setting up
setup(
    # the name must match the folder name 'verysimplemodule'
    name="everyshillingsproto",
    version=VERSION,
    author="AppsLabKe",
    author_email="",
    description=DESCRIPTION,
    long_description=LONG_DESCRIPTION,
    packages=find_packages(),
    install_requires=[
        'grpcio', 'protobuf'
    ],
    keywords=['python', 'grpc'],
    classifiers=[
        "Programming Language :: Python :: 2",
        "Programming Language :: Python :: 3",
    ]
)
