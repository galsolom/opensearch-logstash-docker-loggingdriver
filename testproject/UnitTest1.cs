using System;
using Newtonsoft.Json;
using FluentAssertions;

using Xunit;
using Xunit.Abstractions;
using Newtonsoft.Json.Linq;
using System.Collections.Generic;
using System.Collections;
using FluentAssertions.Formatting;
using FluentAssertions.Execution;

namespace jtest
{
    public class UnitTest1
    {

        [Theory]
        [InlineData("1")]
        [Trait("tc","1338")]
        public void Test1(string one)
        {
            string actual = "1";
            actual.Should().Be("1");

        }
        [Theory]
        [InlineData("1")]
        [Trait("tc","1337")]
        public void Test2(string one)
        {
            string str = "b";
            str.Should().Be("b", "because I expected list {0} to contain that many items",str);
        }
    }
}