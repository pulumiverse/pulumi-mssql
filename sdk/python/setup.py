# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import errno
from setuptools import setup, find_packages
from setuptools.command.install import install
from subprocess import check_call


VERSION = "0.0.0"
def readme():
    try:
        with open('README.md', encoding='utf-8') as f:
            return f.read()
    except FileNotFoundError:
        return "mssql Pulumi Package - Development Version"


setup(name='pulumiverse_mssql',
      python_requires='>=3.8',
      version=VERSION,
      description="A Pulumi Provider for Microsoft SQL Server and Azure SQL",
      long_description=readme(),
      long_description_content_type='text/markdown',
      keywords='pulumi mssql category/database',
      url='https://github.com/pulumiverse/pulumi-mssql',
      project_urls={
          'Repository': 'https://github.com/pulumiverse/pulumi-mssql'
      },
      license='Apache-2.0',
      packages=find_packages(),
      package_data={
          'pulumiverse_mssql': [
              'py.typed',
              'pulumi-plugin.json',
          ]
      },
      install_requires=[
          'parver>=0.2.1',
          'pulumi>=3.0.0,<4.0.0',
          'semver>=2.8.1'
      ],
      zip_safe=False)
